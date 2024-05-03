export type AuthJwtPayload = {
  sub: string;
  name: string;
  exp: number;
};

export type AuthTokens = {
  access: string;
  refresh: string;
};

export type AuthSignUpResponse = {
  id: string;
};

export type AuthSignInResponse = {
  user: {
    id: string;
    name: string;
  };
  tokens: AuthTokens;
};
