import { doRequest } from "./doRequest";

export interface FetchWordArgs {
  isFibo: boolean;
  limit: number;
  idToken?: string;
}

export interface FetchWordData {
  count: number;
  page: number;
  pages: number;
  limit: number;
  words: Word[];
}

export interface Word {
  id: string;
  userId: string;
  word: string;
  definition: string;
  refUrl: string;
  emailReminder: Boolean;
  startDate: string;
}

export const fetchWords = async (
  key: string,
  args: FetchWordArgs,
  page: number = 1
): Promise<FetchWordData> => {
  const { data, error } = await doRequest<FetchWordData>({
    method: "get",
    url: "/api/words",
    params: {
      page: page,
      limit: args.limit,
      fib: args.isFibo ? "true" : "false",
    },
    headers: {
      Authorization: `Bearer ${args.idToken}`,
    },
  });

  if (error) {
    throw error;
  }

  if (!data) {
    throw new Error("Unknown error");
  }

  return data;
};
