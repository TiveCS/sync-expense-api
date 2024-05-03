import { z } from "zod";

export const AuthSignUpSchema = z.object({
  email: z.string().email(),
  name: z.string().min(1),
  password: z.string().min(6),
});

export const AuthSignInSchema = z.object({
  email: z.string().email(),
  password: z.string().min(1),
});

export type AuthSignUpDTO = z.infer<typeof AuthSignUpSchema>;
export type AuthSignInDTO = z.infer<typeof AuthSignInSchema>;
