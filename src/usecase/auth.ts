import { AuthExceptions } from "@/exceptions/auth";
import { AppException } from "@/exceptions/common";
import { createTokenExpireTime } from "@/helpers/jwt";
import { prisma } from "@/lib/prisma";
import { AuthSignInDTO, AuthSignUpDTO } from "@/schemas/auth";
import {
  AuthJwtPayload,
  AuthSignInResponse,
  AuthSignUpResponse,
} from "@/types/auth";
import { hash, verify } from "argon2";
import * as jwt from "hono/jwt";

export async function signUpUsecase(
  dto: AuthSignUpDTO
): Promise<AuthSignUpResponse> {
  const existsUser = await prisma.user.findUnique({
    where: { email: dto.email },
  });

  if (existsUser) throw new AppException(AuthExceptions.USER_EXISTS);

  const hashedPassword = await hash(dto.password);

  const user = await prisma.user.create({
    data: {
      email: dto.email,
      name: dto.name,
      password: hashedPassword,
    },
    select: { id: true },
  });

  return {
    id: user.id,
  };
}

export async function signInUsecase(
  dto: AuthSignInDTO,
  secrets: { access: string; refresh: string }
): Promise<AuthSignInResponse> {
  const foundUser = await prisma.user.findUnique({
    where: { email: dto.email },
  });

  if (!foundUser) throw new AppException(AuthExceptions.INVALID_CREDENTIALS);

  const isPasswordValid = await verify(foundUser.password, dto.password);

  if (!isPasswordValid)
    throw new AppException(AuthExceptions.INVALID_CREDENTIALS);

  const jwtPayload: Omit<AuthJwtPayload, "exp"> = {
    sub: foundUser.id,
    name: foundUser.name,
  };

  const [accessToken, refreshToken] = await Promise.all([
    jwt.sign({ ...jwtPayload, exp: createTokenExpireTime(15) }, secrets.access),
    jwt.sign(
      { ...jwtPayload, exp: createTokenExpireTime(60 * 60 * 7) },
      secrets.refresh
    ),
  ]);

  return {
    user: { id: foundUser.id, name: foundUser.name },
    tokens: { access: accessToken, refresh: refreshToken },
  };
}

export async function refreshTokenUsecase(
  payload: AuthJwtPayload,
  secrets: { access: string; refresh: string }
) {
  const foundUser = await prisma.user.findUnique({
    where: { id: payload.sub },
  });

  if (!foundUser) throw new AppException(AuthExceptions.INVALID_TOKEN);

  const jwtPayload: Omit<AuthJwtPayload, "exp"> = {
    sub: foundUser.id,
    name: foundUser.name,
  };

  const accessToken = await jwt.sign(
    { ...jwtPayload, exp: createTokenExpireTime(15) },
    secrets.access
  );
  const newRefreshToken = await jwt.sign(
    { ...jwtPayload, exp: createTokenExpireTime(60 * 60 * 7) },
    secrets.refresh
  );

  return { access: accessToken, refresh: newRefreshToken };
}

export async function profileUsecase(payload: AuthJwtPayload) {
  return payload;
}
