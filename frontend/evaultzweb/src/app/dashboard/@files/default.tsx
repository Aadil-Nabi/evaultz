import Link from "next/link";

export default function Files() {
  return (
    <>
      <div>
        <div>All Files</div>
        <ul>
          <li>file a</li>
          <li>file b</li>
        </ul>
        <Link href={"/myfiles"}>My Documents</Link>
      </div>
    </>
  );
}
