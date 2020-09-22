import React from "react";
import { useQuery } from "react-query";
import { FetchWordData, fetchWords } from "../data/fetchWords";
import { useAuth } from "../store/auth";

const Overview: React.FC = () => {
  const idToken = useAuth((state) => state.idToken);

  const { isLoading, isError, data, error } = useQuery<FetchWordData, Error>(
    ["words", { isFibo: true, page: 1, limit: 10, idToken }],
    fetchWords
  );

  const wordList = data && data.words.map((word) => <p>{word.word}</p>);
  return (
    <>
      <h1 className="title is-3">Today's Words</h1>

      {isLoading && <p>Loading</p>}
      {isError && <p>{error?.message}</p>}
      {wordList}
    </>
  );
};

export default Overview;
