import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export default function proxy(req: NextRequest) {
  const token = req.cookies.get("jwt")?.value;

  const isAuthPage =
    req.nextUrl.pathname.startsWith("/signin") ||
    req.nextUrl.pathname.startsWith("/signup");

  // 🚫 If visiting protected route without token → redirect
  if (!token && !isAuthPage) {
    return NextResponse.redirect(new URL("/signin", req.url));
  }

  // 🔄 If user already logged in & tries to visit signin/signup → redirect
  if (token && isAuthPage) {
    return NextResponse.redirect(new URL("/dashboard", req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: [
    "/dashboard/:path*", // protect all dashboard routes
    "/files/:path*", // protect user files
    "/profile/:path*", // etc.
    // "/signin",
    // "/signup",
  ],
};
