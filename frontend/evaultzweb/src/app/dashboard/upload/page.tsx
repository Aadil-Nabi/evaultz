// app/dashboard/upload/page.tsx
"use client";

import { useRef, useState } from "react";

export default function UploadPage() {
  const fileRef = useRef<HTMLInputElement | null>(null);
  const [selected, setSelected] = useState<File | null>(null);
  const [loading, setLoading] = useState(false);

  const onPick = () => fileRef.current?.click();
  const onFile = (e: React.ChangeEvent<HTMLInputElement>) => {
    const f = e.target.files?.[0] ?? null;
    setSelected(f);
  };

  const upload = async () => {
    if (!selected) return;
    setLoading(true);
    // TODO: implement actual upload to your backend / signed URL flow
    await new Promise((r) => setTimeout(r, 1200));
    setLoading(false);
    alert("Uploaded (demo)");
  };

  return (
    <div className="space-y-6">
      <div className="bg-white rounded-lg p-6 shadow-sm">
        <h3 className="text-lg font-semibold">
          Upload Files to your Personal storage
        </h3>
        <p className="text-sm text-gray-500">
          Upload files here or create a folder
        </p>

        <div className="mt-6 flex gap-3">
          <button
            onClick={() => alert("Create folder demo")}
            className="px-4 py-2 rounded-md bg-black text-white"
          >
            Create Folder
          </button>
          <button onClick={onPick} className="px-4 py-2 rounded-md border">
            Upload File
          </button>
          <input
            ref={fileRef}
            className="hidden"
            type="file"
            onChange={onFile}
          />
        </div>

        {selected && (
          <div className="mt-4">
            <div className="text-sm">
              Selected: {selected.name} (
              {(selected.size / 1024 / 1024).toFixed(2)} MB)
            </div>
            <div className="mt-2 flex gap-2">
              <button
                onClick={upload}
                className="px-4 py-2 rounded-md bg-indigo-600 text-white"
                disabled={loading}
              >
                {loading ? "Uploading..." : "Start upload"}
              </button>
              <button
                onClick={() => setSelected(null)}
                className="px-4 py-2 rounded-md border"
              >
                Cancel
              </button>
            </div>
          </div>
        )}
      </div>

      <div className="bg-white rounded-lg p-6 shadow-sm">
        <h4 className="text-sm font-semibold mb-2">Upload history</h4>
        <div className="h-40 bg-gray-50 rounded" />
      </div>
    </div>
  );
}
