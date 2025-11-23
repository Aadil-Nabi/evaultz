import axiosClient from "../axiosClient";

export async function SignOutUser(): Promise<void> {
  const response = await axiosClient.post("/api/v1/signout");
}
