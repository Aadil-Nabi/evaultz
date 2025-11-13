"use client";

import { useState } from "react";
import { StatusLoader } from "@/components/internal/statusLoader";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import {
  Field,
  FieldDescription,
  FieldGroup,
  FieldLabel,
  FieldLegend,
  FieldSeparator,
  FieldSet,
} from "@/components/ui/field";
import { Checkbox } from "@/components/ui/checkbox";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { SuccessSooner } from "@/components/internal/successToast";
import { redirect } from "next/navigation";

export default function UploadFilePage() {
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // Show loader
    setIsSubmitting(true);

    // Simulate upload or API call
    await new Promise((resolve) => setTimeout(resolve, 3000));

    // Hide loader or navigate away

    setIsSubmitting(false);
    <SuccessSooner />;
    redirect("/");
  };

  if (isSubmitting) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-50">
        <div className="text-center">
          <StatusLoader />
          <p className="mt-4 text-gray-600 font-medium">
            Uploading your file, please wait...
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-50 p-6">
      <div className="w-full max-w-md bg-white p-8 rounded-2xl shadow-md space-y-6">
        <form className="space-y-6" onSubmit={handleSubmit}>
          <FieldGroup>
            <FieldSet>
              <FieldLegend className="text-xl font-semibold text-center">
                Please Upload Your File
              </FieldLegend>
              <FieldDescription className="text-center text-gray-500">
                All big files of any type can be uploaded here.
              </FieldDescription>
            </FieldSet>

            <FieldSeparator />

            <FieldSet>
              <FieldLegend>Share file with others</FieldLegend>
              <FieldDescription>
                Check this box if you want others to view this file as well.
              </FieldDescription>
              <FieldGroup>
                <Field orientation="horizontal">
                  <Checkbox id="share-checkbox" defaultChecked />
                  <FieldLabel
                    htmlFor="share-checkbox"
                    className="font-normal ml-2"
                  >
                    Share
                  </FieldLabel>
                </Field>
              </FieldGroup>
            </FieldSet>

            <FieldSet>
              <FieldGroup>
                <Field>
                  <FieldLabel htmlFor="file-description">
                    Description (optional)
                  </FieldLabel>
                  <Textarea
                    id="file-description"
                    placeholder="Add any additional comments"
                    className="resize-none"
                  />
                </Field>
              </FieldGroup>
            </FieldSet>

            <div className="grid w-full gap-3">
              <Label htmlFor="file">Please select the file</Label>
              <Input id="file" type="file" />
            </div>

            <Field
              orientation="horizontal"
              className="justify-center space-x-3"
            >
              <Button type="submit">Submit</Button>
              <Button variant="outline" type="button">
                Cancel
              </Button>
            </Field>
          </FieldGroup>
        </form>
      </div>
    </div>
  );
}
