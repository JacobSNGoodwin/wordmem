import React from "react";
import { useQuery } from "react-query";
import { FetchWordData, fetchWords } from "../data/fetchWords";
import { useAuth } from "../store/auth";

const Edit: React.FC = () => {
  const idToken = useAuth((state) => state.idToken);

  const { isLoading, isError, data, error } = useQuery<FetchWordData, Error>(
    ["words", { isFibo: false, page: 1, limit: 10, idToken }],
    fetchWords
  );
  console.log(data?.words);
  return (
    <>
      <h1 className="title is-3">Your Word List</h1>

      {isLoading && <p>Loading</p>}
      {isError && <p>{error?.message}</p>}
      {data && <p>{data.words}</p>}
    </>
  );
};

export default Edit;
