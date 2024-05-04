import { TransactionCategory } from "@prisma/client";

export type CreateAccountResponse = {
  id: string;
};

export type UpdateAccountResponse = {
  id: string;
};

export type AccountOverview = {
  id: string;
  name: string;
  createdAt: Date;
};

export type AccountDetails = {
  id: string;
  name: string;
  _count: {
    transactions: number;
  };
};

export type GetAccountsResponse = AccountOverview[];

export type GetAccountDetailsResponse = AccountDetails;
