import axiosClient from "../axiosClient";

interface SignInPayload {
  email: string;
  password: string;
  tenant: string;
}

export async function signInUser(data: SignInPayload): Promise<any> {
  try {
    const response = await axiosClient.post("/api/v1/signin", data);
    return response;
  } catch (error) {
    throw error;
  }
}
