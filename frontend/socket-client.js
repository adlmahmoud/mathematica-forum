// ===================== SOCKET.IO CLIENT =====================
// Installe socket.io-client si besoin :
// npm install socket.io-client
// ou via CDN dans HTML

import { io } from "socket.io-client";

// Connexion au serveur
const socket = io("http://localhost:3000");

// ===================== ETAT UTILISATEUR =====================
let currentUser = null;
let currentChatUser = null;

// ===================== AUTH INIT =====================
export function initSocket(user) {
  currentUser = user;

  socket.on("connect", () => {
    console.log("Connecté au chat :", socket.id);
  });

  socket.on("disconnect", () => {
    console.log("Déconnecté du chat");
  });

  // Réception des messages DM
  socket.on("dm", (data) => {
    console.log("DM reçu :", data);

    if (data.to === currentUser.id || data.from === currentUser.id) {
      displayMessage(data);
    }
  });
}

// ===================== ENVOI MESSAGE DM =====================
export function sendDM(toUserId, message) {
  if (!message.trim()) return;

  const payload = {
    from: currentUser.id,
    to: toUserId,
    message,
    timestamp: new Date(),
  };

  socket.emit("dm", payload);

  // affichage local immédiat
  displayMessage(payload, true);
}

// ===================== CHOISIR UTILISATEUR CHAT =====================
export function openChat(userId) {
  currentChatUser = userId;
}

// ===================== AFFICHAGE MESSAGE =====================
function displayMessage(data, isMine = false) {
  const chatBox = document.getElementById("chat-box");

  if (!chatBox) return;

  const msg = document.createElement("div");
  msg.className = isMine ? "msg mine" : "msg";

  msg.innerHTML = `
    <div class="msg-content">
      <strong>${isMine ? "Moi" : data.from}</strong>
      <p>${data.message}</p>
      <small>${new Date(data.timestamp).toLocaleTimeString()}</small>
    </div>
  `;

  chatBox.appendChild(msg);
  chatBox.scrollTop = chatBox.scrollHeight;
}

// ===================== UI SEND BUTTON =====================
export function bindChatInput(inputId, buttonId) {
  const input = document.getElementById(inputId);
  const button = document.getElementById(buttonId);

  if (!input || !button) return;

  button.addEventListener("click", () => {
    if (currentChatUser) {
      sendDM(currentChatUser, input.value);
      input.value = "";
    }
  });

  input.addEventListener("keypress", (e) => {
    if (e.key === "Enter") {
      button.click();
    }
  });
}

// ===================== NOTIFICATIONS =====================
socket.on("notification", (data) => {
  console.log("Notification :", data);
  alert(data.message);
});

export default socket;