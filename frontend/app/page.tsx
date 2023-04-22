import Image from "next/image";
import styles from "./page.module.css";

async function getData() {
  const res = await fetch("http://localhost:8080/api/v1/entries");
  // The return value is *not* serialized
  // You can return Date, Map, Set, etc.

  // Recommendation: handle errors
  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error("Failed to fetch data");
  }

  return res.json();
}

export default async function Home() {
  const data = await getData();
  console.log(data);

  return (
    <main className={styles.main}>
      <div className={styles.description}>
        <h1>Hello World</h1>
      </div>
    </main>
  );
}
