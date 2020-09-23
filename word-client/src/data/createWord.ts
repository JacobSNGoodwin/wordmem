import { doRequest } from "./doRequest";
import { Word } from "./fetchWords";

type CreateWordData = Word;
type CreateWordInput = {
  idToken: string;
  word: string;
  definition: string;
  refUrl: string;
};

const createWord = async ({
  idToken,
  word,
  definition,
  refUrl,
}: CreateWordInput) => {
  const { data, error } = await doRequest<CreateWordData>({
    method: "post",
    url: "/api/words",
    headers: {
      Authorization: `Bearer ${idToken}`,
    },
    data: {
      word,
      definition,
      refUrl,
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

export default createWord;
