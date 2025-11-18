"use client";

import { useRef } from "react";

import { IconFolderCode } from "@tabler/icons-react";

import { Button } from "@/components/ui/button";
import {
  Empty,
  EmptyContent,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from "@/components/ui/empty";
import { Input } from "../ui/input";

export default function UploadFilePage() {
  const fileRef = useRef<HTMLInputElement>(null);

  const handleUploadClick = () => {
    fileRef.current?.click();
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    console.log("Selected file:", file);
    // upload logic here...
  };

  return (
    <Empty>
      <EmptyHeader>
        <EmptyMedia variant="icon">
          <IconFolderCode />
        </EmptyMedia>
        <EmptyTitle>Upload Files to your Personal storage</EmptyTitle>
        <EmptyDescription>
          Upload files here or create a folder
        </EmptyDescription>
      </EmptyHeader>
      <EmptyContent>
        <div className="flex gap-2">
          <Button>Create Folder</Button>
          <Button onClick={handleUploadClick} variant="outline">
            Upload File
          </Button>
        </div>

        <Input
          id="file"
          type="file"
          ref={fileRef}
          onChange={handleFileChange}
          className="hidden"
        />
      </EmptyContent>
    </Empty>
  );
}
