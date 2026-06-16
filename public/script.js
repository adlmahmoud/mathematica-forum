
const API = "http://localhost:3000/api";

let token = localStorage.getItem("token");
let currentUser = null;
let questionsCache = [];

const questionsList = document.getElementById("questions-list");
const questionTemplate = document.getElementById("question-template");

// ===================== INIT =====================
document.addEventListener("DOMContentLoaded", () => {
  initNavigation();
  initTheme();
  initQuestionForm();
  initSearch();
  loadQuestions();
  loadLeaderboard();
  loadProfile();
  initPremium();
});

// ===================== NAVIGATION =====================
function initNavigation() {
  document.querySelectorAll(".nav-btn").forEach(btn => {
    btn.addEventListener("click", () => {
      document.querySelectorAll(".nav-btn").forEach(el =>
        el.classList.remove("active")
      );

      document.querySelectorAll(".panel").forEach(el => {
        el.classList.remove("active-panel");
        el.classList.add("hidden");
      });

      btn.classList.add("active");

      const target = document.getElementById(btn.dataset.panel);
      target.classList.remove("hidden");
      target.classList.add("active-panel");
    });
  });
}

// ===================== THEME =====================
function initTheme() {
  const button = document.getElementById("theme-toggle");
  const saved = localStorage.getItem("theme");

  if (saved === "light") {
    document.body.classList.add("light");
  }

  button.addEventListener("click", () => {
    document.body.classList.toggle("light");

    localStorage.setItem(
      "theme",
      document.body.classList.contains("light") ? "light" : "dark"
    );
  });
}

// ===================== CREATE QUESTION =====================
function initQuestionForm() {
  document.getElementById("new-question-btn")
    .addEventListener("click", () => {
      document.getElementById("question-form").classList.toggle("hidden");
    });

  document.getElementById("submit-question")
    .addEventListener("click", createQuestion);
}

async function createQuestion() {
  if (!token) return alert("Connexion requise");

  const title = document.getElementById("question-title").value;
  const body = document.getElementById("question-body").value;
  const tags = document.getElementById("question-tags").value
    .split(",")
    .map(t => t.trim());

  if (!title || !body) return alert("Champs requis");

  const res = await fetch(`${API}/questions`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify({
      title,
      content: body,
      tags
    })
  });

  if (res.ok) {
    document.getElementById("question-title").value = "";
    document.getElementById("question-body").value = "";
    document.getElementById("question-tags").value = "";
    loadQuestions();
  }
}

// ===================== LOAD QUESTIONS =====================
async function loadQuestions() {
  const res = await fetch(`${API}/questions`);
  questionsCache = await res.json();
  renderQuestions(questionsCache);
}

// ===================== RENDER =====================
function renderQuestions(questions) {
  questionsList.innerHTML = "";

  questions.forEach(q => {
    const node = questionTemplate.content.cloneNode(true);

    node.querySelector(".question-title").textContent = q.title;
    node.querySelector(".question-body").textContent = q.content;
    node.querySelector(".vote-count").textContent = q.votes;

    node.querySelector(".question-author").textContent =
      q.author?.username || "Utilisateur";

    node.querySelector(".question-date").textContent =
      new Date(q.createdAt).toLocaleDateString("fr-FR");

    const tagsContainer = node.querySelector(".question-tags");

    (q.tags || []).forEach(tag => {
      const span = document.createElement("span");
      span.textContent = tag;
      tagsContainer.appendChild(span);
    });

    node.querySelector(".vote-up").addEventListener("click", () =>
      voteQuestion(q._id, 1)
    );

    node.querySelector(".vote-down").addEventListener("click", () =>
      voteQuestion(q._id, -1)
    );

    const answerBtn = node.querySelector(".answer-btn");
    const answerInput = node.querySelector(".answer-input");

    answerBtn.addEventListener("click", () =>
      answerQuestion(q._id, answerInput.value)
    );

    const answersContainer = node.querySelector(".answers-container");

    (q.answers || []).forEach(a => {
      const div = document.createElement("div");
      div.className = "answer-item";
      div.innerHTML = `<p>${a.content}</p>`;
      answersContainer.appendChild(div);
    });

    questionsList.appendChild(node);
  });
}

// ===================== VOTE =====================
async function voteQuestion(id, value) {
  if (!token) return alert("Connexion requise");

  await fetch(`${API}/questions/${id}/vote`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify({ value })
  });

  loadQuestions();
}

// ===================== ANSWER =====================
async function answerQuestion(id, content) {
  if (!content) return;
  if (!token) return alert("Connexion requise");

  await fetch(`${API}/questions/${id}/answer`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify({ content })
  });

  loadQuestions();
}

// ===================== SEARCH =====================
function initSearch() {
  document.getElementById("search-input").addEventListener("input", e => {
    const value = e.target.value.toLowerCase();

    const filtered = questionsCache.filter(q =>
      q.title.toLowerCase().includes(value) ||
      q.content.toLowerCase().includes(value)
    );

    renderQuestions(filtered);
  });
}

// ===================== LEADERBOARD =====================
async function loadLeaderboard() {
  const container = document.getElementById("leaderboard-list");
  if (!container) return;

  const res = await fetch(`${API}/users/leaderboard`);
  const users = await res.json();

  container.innerHTML = "";

  users.forEach((u, i) => {
    const div = document.createElement("div");
    div.className = "leaderboard-item";

    div.innerHTML = `
      <span>#${i + 1} ${u.username}</span>
      <strong>${u.points} XP</strong>
    `;

    container.appendChild(div);
  });
}

// ===================== PROFILE =====================
async function loadProfile() {
  if (!token) return;

  try {
    const res = await fetch(`${API}/users/leaderboard`, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    });

    const data = await res.json();
    currentUser = data[0]; // simplifié

  } catch (err) {
    console.error(err);
  }
}

// ===================== PREMIUM =====================
function initPremium() {
  const btn = document.getElementById("buy-premium");
  if (!btn) return;

  btn.addEventListener("click", activatePremium);
}

async function activatePremium() {
  if (!token) return alert("Connexion requise");

  await fetch(`${API}/premium/activate`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`
    }
  });

  alert("Premium activé !");
  loadProfile();
}