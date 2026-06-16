const mongoose = require("mongoose");

const userSchema = new mongoose.Schema(
  {
    username: {
      type: String,
      required: true,
      unique: true,
      trim: true
    },

    email: {
      type: String,
      required: true,
      unique: true,
      lowercase: true
    },

    password: {
      type: String,
      required: true
    },

    role: {
      type: String,
      enum: [
        "Étudiant",
        "Professeur",
        "Ingénieur",
        "Admin",
        "Modérateur"
      ],
      default: "Étudiant"
    },

    xp: {
      type: Number,
      default: 0
    },

    level: {
      type: Number,
      default: 1
    },

    premium: {
      type: Boolean,
      default: false
    },

    avatar: {
      type: String,
      default: ""
    },

    badges: [
      {
        name: String,
        description: String
      }
    ]
  },
  {
    timestamps: true
  }
);

module.exports = mongoose.model(
  "User",
  userSchema
);