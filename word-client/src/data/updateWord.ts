import { doRequest } from "./doRequest";
import { Word } from "./fetchWords";

type UpdateWordData = Word;
type UpdateWordInput = {
  idToken?: string;
  id?: string;
  word: string;
  definition: string;
  refUrl: string;
  startDate?: string;
};

// handles both create and update data
const updateWord = async ({
  idToken,
  id,
  word,
  definition,
  refUrl,
  startDate,
}: UpdateWordInput) => {
  const url = id ? `/api/words/${id}` : `/api/words`;
  const method = id ? "put" : "post";
  const bodyData = id
    ? {
        word,
        definition,
        refUrl,
        startDate,
      }
    : { word, definition, refUrl };

  const { data, error } = await doRequest<UpdateWordData>({
    method,
    url,
    headers: {
      Authorization: `Bearer ${idToken}`,
    },
    data: bodyData,
  });

  if (error) {
    throw error;
  }

  if (!data) {
    throw new Error("Unknown error");
  }

  return data;
};

export default updateWord;
