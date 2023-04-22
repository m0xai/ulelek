import { log } from "console";

async function getEntry(id: Number) {
  const res = await fetch("http://localhost:8080/api/v1/entries/" + id);
  // The return value is *not* serialized
  // You can return Date, Map, Set, etc.

  // Recommendation: handle errors
  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error("Failed to fetch data");
  }

  return res.json();
}

export default async function Entry({ params }: any) {
  const entry = await getEntry(params.id);
  log(entry);
  return <div>My Post {entry.content} </div>;
}
