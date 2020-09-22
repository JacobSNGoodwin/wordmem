import React from "react";
import { useQuery } from "react-query";
import Loader from "../components/ui/Loader";
import WordList from "../components/WordList";
import { FetchWordData, fetchWords } from "../data/fetchWords";
import { useAuth } from "../store/auth";

const Edit: React.FC = () => {
  const idToken = useAuth((state) => state.idToken);

  const { isLoading, isError, data, error } = useQuery<FetchWordData, Error>(
    ["words", { isFibo: false, page: 1, limit: 10, idToken }],
    fetchWords,
    { staleTime: 3000 }
  );

  const wordList = data?.words ? <WordList words={data.words} /> : undefined;

  return (
    <>
      <h1 className="title is-3">Your Word List</h1>
      <div className="buttons is-centered">
        <button className="button is-info">Create Word</button>
      </div>

      {isLoading && <Loader radius={200} />}
      {isError && <p>{error?.message}</p>}
      {wordList}
    </>
  );
};

export default Edit;
