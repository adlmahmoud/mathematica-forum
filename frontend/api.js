
// ===================== CONFIG =====================
const API_URL = "http://localhost:3000/api";

// ===================== STORAGE TOKEN =====================
function setToken(token) {
  localStorage.setItem("token", token);
}

function getToken() {
  return localStorage.getItem("token");
}

function removeToken() {
  localStorage.removeItem("token");
}

// ===================== HEADERS =====================
function authHeaders() {
  return {
    "Content-Type": "application/json",
    Authorization: "Bearer " + getToken(),
  };
}

// ===================== FETCH WRAPPER =====================
async function request(url, options = {}) {
  const res = await fetch(API_URL + url, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...(options.headers || {}),
      ...(getToken() ? { Authorization: "Bearer " + getToken() } : {}),
    },
  });

  const data = await res.json();
  if (!res.ok) throw data;

  return data;
}

// ===================== AUTH =====================

// REGISTER
export async function register(username, email, password) {
  return await request("/auth/register", {
    method: "POST",
    body: JSON.stringify({ username, email, password }),
  });
}

// LOGIN
export async function login(email, password) {
  const data = await request("/auth/login", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });

  if (data.token) {
    setToken(data.token);
  }

  return data;
}

// LOGOUT
export function logout() {
  removeToken();
}

// GET CURRENT USER (via decode simple côté backend)
export async function getMe() {
  return await request("/users/leaderboard"); // fallback simple
}

// ===================== QUESTIONS =====================

// GET ALL QUESTIONS
export async function getQuestions() {
  return await request("/questions");
}

// GET ONE QUESTION
export async function getQuestion(id) {
  return await request(`/questions/${id}`);
}

// CREATE QUESTION
export async function createQuestion(title, content) {
  return await request("/questions", {
    method: "POST",
    body: JSON.stringify({ title, content }),
  });
}

// DELETE QUESTION
export async function deleteQuestion(id) {
  return await request(`/questions/${id}`, {
    method: "DELETE",
  });
}

// EDIT QUESTION
export async function editQuestion(id, title, content) {
  return await request(`/questions/${id}`, {
    method: "PUT",
    body: JSON.stringify({ title, content }),
  });
}

// ===================== ANSWERS =====================

// ADD ANSWER
export async function addAnswer(questionId, content) {
  return await request(`/questions/${questionId}/answer`, {
    method: "POST",
    body: JSON.stringify({ content }),
  });
}

// ACCEPT ANSWER
export async function acceptAnswer(qid, aid) {
  return await request(`/questions/${qid}/accept/${aid}`, {
    method: "POST",
  });
}

// ===================== VOTES =====================

// VOTE QUESTION (+1 / -1)
export async function voteQuestion(id, value) {
  return await request(`/questions/${id}/vote`, {
    method: "POST",
    body: JSON.stringify({ value }),
  });
}

// ===================== SEARCH =====================
export async function searchQuestions(query) {
  return await request(`/search?q=${encodeURIComponent(query)}`);
}

// ===================== PREMIUM =====================
export async function activatePremium() {
  return await request("/premium/activate", {
    method: "POST",
  });
}

// ===================== USERS =====================

// LEADERBOARD
export async function getLeaderboard() {
  return await request("/users/leaderboard");
}

// UPLOAD AVATAR
export async function uploadAvatar(file) {
  const formData = new FormData();
  formData.append("avatar", file);

  const res = await fetch(API_URL + "/users/avatar", {
    method: "POST",
    headers: {
      Authorization: "Bearer " + getToken(),
    },
    body: formData,
  });

  return await res.json();
}

// ===================== HELPERS =====================
export function isLogged() {
  return !!getToken();
}

export function getTokenValue() {
  return getToken();
}