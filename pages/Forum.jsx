import { useEffect, useState } from "react";
import API from "../api/apiClient";

export default function Forum() {
  const [questions, setQuestions] = useState([]);

  useEffect(() => {
    API.get("/questions").then((res) => {
      setQuestions(res.data);
    });
  }, []);

  return (
    <div>
      <h1>Forum Math+</h1>

      {questions.map((q) => (
        <div key={q._id}>
          <h3>{q.title}</h3>
          <p>{q.content}</p>
        </div>
      ))}
    </div>
  );
}