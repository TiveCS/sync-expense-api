import { AccountsExceptions } from "@/exceptions/accounts";
import { AppException } from "@/exceptions/common";
import { prisma } from "@/lib/prisma";
import { CreateAccountDTO, UpdateAccountDTO } from "@/schemas/accounts";
import {
  CreateAccountResponse,
  GetAccountDetailsResponse,
  GetAccountsResponse,
  UpdateAccountResponse,
} from "@/types/accounts";
import { Prisma } from "@prisma/client";

export async function createAccountUsecase(
  userId: string,
  dto: CreateAccountDTO
): Promise<CreateAccountResponse> {
  const account = await prisma.account.create({
    data: {
      name: dto.name,
      user: { connect: { id: userId } },
    },
    select: { id: true },
  });

  return { id: account.id };
}

export async function updateAccountUsecase(
  accountId: string,
  dto: UpdateAccountDTO
): Promise<UpdateAccountResponse> {
  try {
    const account = await prisma.account.update({
      where: { id: accountId },
      data: { name: dto.name },
      select: { id: true },
    });

    return { id: account.id };
  } catch (error) {
    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      if (error.code === "P2025") {
        throw new AppException(AccountsExceptions.ACCOUNT_NOT_FOUND);
      }
    }
    throw error;
  }
}

export async function deleteAccountUsecase(accountId: string): Promise<void> {
  try {
    await prisma.account.delete({
      where: { id: accountId },
      select: { id: true },
    });
  } catch (error) {
    if (error instanceof Prisma.PrismaClientKnownRequestError) {
      if (error.code === "P2025") {
        throw new AppException(AccountsExceptions.ACCOUNT_NOT_FOUND);
      }
    }
    throw error;
  }
}

export async function getAccountsUsecase(
  userId: string
): Promise<GetAccountsResponse> {
  const accounts = await prisma.account.findMany({
    where: { userId },
  });

  return accounts.map((account) => ({
    id: account.id,
    name: account.name,
    createdAt: account.createdAt,
  }));
}

export async function getAccountDetailsUsecase(
  accountId: string
): Promise<GetAccountDetailsResponse> {
  const account = await prisma.account.findUnique({
    where: { id: accountId },
    include: {
      _count: { select: { transactions: true } },
    },
  });

  if (!account) {
    throw new AppException(AccountsExceptions.ACCOUNT_NOT_FOUND);
  }

  return {
    id: account.id,
    name: account.name,
    _count: { transactions: account._count.transactions },
  };
}
