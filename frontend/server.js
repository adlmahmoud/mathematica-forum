require("dotenv").config();

const express = require("express");
const mongoose = require("mongoose");
const cors = require("cors");
const helmet = require("helmet");
const morgan = require("morgan");
const bcrypt = require("bcrypt");
const jwt = require("jsonwebtoken");
const rateLimit = require("express-rate-limit");
const multer = require("multer");
const http = require("http");
const { Server } = require("socket.io");
const path = require("path");
const fs = require("fs");

// Models
const User = require("./models/User");
const Question = require("./models/Question");

const app = express();
const server = http.createServer(app);
const io = new Server(server, {
  cors: { origin: "*" },
});

// ===================== CONFIG =====================
const PORT = process.env.PORT || 3000;
const MONGO_URI = process.env.MONGO_URI;
const JWT_SECRET = process.env.JWT_SECRET;

// ===================== MIDDLEWARE =====================
app.use(express.json());
app.use(cors());
app.use(helmet());
app.use(morgan("dev"));

app.use("/uploads", express.static(path.join(__dirname, "uploads")));

// Rate limit
app.use(
  rateLimit({
    windowMs: 15 * 60 * 1000,
    max: 200,
  })
);

// ===================== DB =====================
mongoose
  .connect(MONGO_URI)
  .then(() => console.log("MongoDB connecté"))
  .catch((err) => console.log(err));

// ===================== MULTER =====================
const storage = multer.diskStorage({
  destination: (req, file, cb) => {
    cb(null, "uploads/");
  },
  filename: (req, file, cb) => {
    cb(null, Date.now() + "-" + file.originalname);
  },
});
const upload = multer({ storage });

// ===================== AUTH MIDDLEWARE =====================
const auth = (req, res, next) => {
  const token = req.header("Authorization");
  if (!token) return res.status(401).json({ msg: "Accès refusé" });

  try {
    const decoded = jwt.verify(token.replace("Bearer ", ""), JWT_SECRET);
    req.user = decoded;
    next();
  } catch (e) {
    res.status(401).json({ msg: "Token invalide" });
  }
};

// ===================== AUTH ROUTES =====================

// REGISTER
app.post("/api/auth/register", async (req, res) => {
  try {
    const { username, email, password } = req.body;

    const existing = await User.findOne({ email });
    if (existing) return res.status(400).json({ msg: "Utilisateur existe" });

    const hash = await bcrypt.hash(password, 10);

    const user = await User.create({
      username,
      email,
      password: hash,
      premium: false,
      role: "user",
      points: 0,
    });

    res.json(user);
  } catch (err) {
    res.status(500).json(err);
  }
});

// LOGIN
app.post("/api/auth/login", async (req, res) => {
  try {
    const { email, password } = req.body;

    const user = await User.findOne({ email });
    if (!user) return res.status(400).json({ msg: "User not found" });

    const match = await bcrypt.compare(password, user.password);
    if (!match) return res.status(400).json({ msg: "Wrong password" });

    const token = jwt.sign(
      { id: user._id, role: user.role },
      JWT_SECRET,
      { expiresIn: "7d" }
    );

    res.json({ token, user });
  } catch (err) {
    res.status(500).json(err);
  }
});

// ===================== USERS =====================
app.get("/api/users/leaderboard", async (req, res) => {
  const users = await User.find().sort({ points: -1 }).limit(50);
  res.json(users);
});

// Upload avatar
app.post("/api/users/avatar", auth, upload.single("avatar"), async (req, res) => {
  const user = await User.findByIdAndUpdate(
    req.user.id,
    { avatar: req.file.path },
    { new: true }
  );

  res.json(user);
});

// ===================== QUESTIONS =====================

// CREATE QUESTION
app.post("/api/questions", auth, async (req, res) => {
  const question = await Question.create({
    title: req.body.title,
    content: req.body.content,
    author: req.user.id,
    votes: 0,
    answers: [],
    createdAt: new Date(),
  });

  res.json(question);
});

// GET ALL QUESTIONS
app.get("/api/questions", async (req, res) => {
  const q = await Question.find().populate("author");
  res.json(q);
});

// GET ONE
app.get("/api/questions/:id", async (req, res) => {
  const q = await Question.findById(req.params.id).populate("author");
  res.json(q);
});

// DELETE QUESTION
app.delete("/api/questions/:id", auth, async (req, res) => {
  const q = await Question.findById(req.params.id);

  if (q.author.toString() !== req.user.id && req.user.role !== "admin") {
    return res.status(403).json({ msg: "Interdit" });
  }

  await q.deleteOne();
  res.json({ msg: "Deleted" });
});

// EDIT QUESTION
app.put("/api/questions/:id", auth, async (req, res) => {
  const q = await Question.findById(req.params.id);

  if (q.author.toString() !== req.user.id) {
    return res.status(403).json({ msg: "Interdit" });
  }

  q.title = req.body.title || q.title;
  q.content = req.body.content || q.content;

  await q.save();
  res.json(q);
});

// ===================== ANSWERS =====================
app.post("/api/questions/:id/answer", auth, async (req, res) => {
  const q = await Question.findById(req.params.id);

  q.answers.push({
    content: req.body.content,
    author: req.user.id,
    votes: 0,
    accepted: false,
    createdAt: new Date(),
  });

  await q.save();
  res.json(q);
});

// ACCEPT ANSWER
app.post("/api/questions/:qid/accept/:aid", auth, async (req, res) => {
  const q = await Question.findById(req.params.qid);

  if (q.author.toString() !== req.user.id)
    return res.status(403).json({ msg: "Only author can accept" });

  q.answers.forEach((a, i) => {
    q.answers[i].accepted = i == req.params.aid;
  });

  await q.save();
  res.json(q);
});

// ===================== VOTES =====================
app.post("/api/questions/:id/vote", auth, async (req, res) => {
  const q = await Question.findById(req.params.id);

  q.votes += req.body.value; // +1 ou -1
  await q.save();

  const user = await User.findById(q.author);
  user.points += req.body.value;
  await user.save();

  res.json(q);
});

// ===================== SEARCH =====================
app.get("/api/search", async (req, res) => {
  const q = req.query.q;

  const results = await Question.find({
    title: { $regex: q, $options: "i" },
  });

  res.json(results);
});

// ===================== PREMIUM =====================
app.post("/api/premium/activate", auth, async (req, res) => {
  const user = await User.findById(req.user.id);
  user.premium = true;
  await user.save();

  res.json(user);
});

// ===================== SOCKET.IO =====================
io.on("connection", (socket) => {
  console.log("User connected");

  socket.on("dm", ({ to, message }) => {
    io.emit("dm", { to, message });
  });

  socket.on("disconnect", () => {
    console.log("User disconnected");
  });
});

// ===================== ADMIN MODERATION =====================
app.get("/api/admin/users", auth, async (req, res) => {
  if (req.user.role !== "admin")
    return res.status(403).json({ msg: "Admin only" });

  const users = await User.find();
  res.json(users);
});

// ===================== START =====================
server.listen(PORT, () => {
  console.log("Server running on port", PORT);
});