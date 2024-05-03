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

export const AuthTokenPayloadSchema = z.object({
  sub: z.string().cuid(),
  name: z.string(),
  exp: z.number().int().positive(),
});

export const AuthSignUpResponseSchema = z.object({
  id: z.string().cuid(),
});

export const AuthSignInResponseSchema = z.object({
  user: z.object({
    id: z.string().cuid(),
    name: z.string(),
  }),
  tokens: z.object({
    access: z.string(),
    refresh: z.string(),
  }),
});

export type AuthSignUpDTO = z.infer<typeof AuthSignUpSchema>;
export type AuthSignInDTO = z.infer<typeof AuthSignInSchema>;

export type AuthSignUpResponse = z.infer<typeof AuthSignUpResponseSchema>;
export type AuthSignInResponse = z.infer<typeof AuthSignInResponseSchema>;
