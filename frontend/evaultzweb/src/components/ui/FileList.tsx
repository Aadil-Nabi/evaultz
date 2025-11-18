// components/ui/FileList.tsx
export default function FileList({ files }: { files: any[] }) {
  return (
    <div className="bg-white rounded-lg p-4 shadow-sm">
      <table className="w-full text-sm">
        <thead>
          <tr className="text-left text-xs text-gray-400">
            <th className="py-2">Name</th>
            <th className="py-2">Size</th>
            <th className="py-2">Updated</th>
            <th className="py-2">Actions</th>
          </tr>
        </thead>
        <tbody className="divide-y">
          {files.map((f) => (
            <tr key={f.id}>
              <td className="py-3">{f.name}</td>
              <td className="py-3 text-gray-500">{f.size}</td>
              <td className="py-3 text-gray-500">{f.updated}</td>
              <td className="py-3">
                <div className="flex gap-2">
                  <button className="text-indigo-600 text-sm">Open</button>
                  <button className="text-gray-600 text-sm">Download</button>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
