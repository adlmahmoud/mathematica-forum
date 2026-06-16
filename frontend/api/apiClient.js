import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:3000/api",
  withCredentials: true,
});

// ===================== TOKEN =====================
let accessToken = null;

export function setToken(token) {
  accessToken = token;
}

export function getToken() {
  return accessToken;
}

// ===================== REQUEST INTERCEPTOR =====================
API.interceptors.request.use((config) => {
  if (accessToken) {
    config.headers.Authorization = `Bearer ${accessToken}`;
  }
  return config;
});

// ===================== RESPONSE INTERCEPTOR =====================
API.interceptors.response.use(
  (res) => res,
  async (err) => {
    const original = err.config;

    if (err.response?.status === 401 && !original._retry) {
      original._retry = true;

      try {
        const { data } = await API.post("/auth/refresh");
        setToken(data.token);

        original.headers.Authorization = `Bearer ${data.token}`;
        return API(original);
      } catch {
        logout();
        window.location.href = "/login";
      }
    }

    return Promise.reject(err);
  }
);

// ===================== AUTH =====================
export async function login(email, password) {
  const { data } = await API.post("/auth/login", { email, password });
  setToken(data.token);
  return data;
}

export async function logout() {
  await API.post("/auth/logout");
  setToken(null);
}

// ===================== QUESTIONS =====================
export const getQuestions = () => API.get("/questions");
export const createQuestion = (data) => API.post("/questions", data);
export const voteQuestion = (id, value) =>
  API.post(`/questions/${id}/vote`, { value });

export default API;