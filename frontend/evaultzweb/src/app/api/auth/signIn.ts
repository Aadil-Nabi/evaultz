import axiosClient from "../axiosClient";

interface SignInPayload {
  email: string;
  password: string;
  companyname: string;
}

export async function signInUser(data: SignInPayload): Promise<any> {
  try {
    const response = await axiosClient.post("/api/v1/signin", data);
    console.log("Signin API Response:", response);
    return response;
  } catch (error) {
    throw error;
  }
}
