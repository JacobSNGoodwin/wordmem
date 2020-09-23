import { doRequest } from "./doRequest";

type DeleteWordData = {
  ids: string[];
};
type DeleteWordInput = {
  idToken?: string;
  id: string;
};

const deleteWord = async ({ idToken, id }: DeleteWordInput) => {
  const { data, error } = await doRequest<DeleteWordData>({
    method: "post",
    url: `/api/words/delete`,
    headers: {
      Authorization: `Bearer ${idToken}`,
    },
    data: {
      wordIds: [id],
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

export default deleteWord;
