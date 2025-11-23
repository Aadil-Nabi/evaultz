import axiosClient from "../axiosClient";

export interface SignUpResponse {
  success: boolean;
  message: string;
}

interface SignUpPayload {
  email: string;
  password: string;
  username: string;
  companyname: string;
}

export async function signUpUser(data: SignUpPayload): Promise<SignUpResponse> {
  try {
    const response = await axiosClient.post("/api/v1/signup", data);
    return response.data;
  } catch (err) {
    console.error("Signup API Error:", err);
    throw err;
  }
}
