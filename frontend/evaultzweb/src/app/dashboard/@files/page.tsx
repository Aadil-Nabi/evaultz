import Link from "next/link";

export default function Files() {
  return (
    <>
      <div>
        <div>
          <h1 className="b">ALL Files</h1>
        </div>
        <ol>
          <li>file a</li>
          <li>file b</li>
        </ol>
        <Link
          href={"/dashboard/myfiles"}
          className="bg-black text-amber-50 p-2  rounded-2xl mt-2"
        >
          My files
        </Link>
      </div>
    </>
  );
}
