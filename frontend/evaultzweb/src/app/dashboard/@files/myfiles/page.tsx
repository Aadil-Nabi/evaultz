import Link from "next/link";

export default function MyFiles() {
  return (
    <>
      <div>
        <div>
          <h1 className="b">My Files</h1>
        </div>
        <ol>
          <li>My doc A</li>
          <li>My Doc B</li>
        </ol>
        <Link
          href={"/dashboard"}
          className="bg-black text-amber-50 p-2 rounded-2xl mt-2"
        >
          All files
        </Link>
      </div>
    </>
  );
}
