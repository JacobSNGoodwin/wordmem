import React, { useState } from "react";
import { useQuery } from "react-query";
import EditWordForm from "../components/EditWordForm";
import Loader from "../components/ui/Loader";
import WordList from "../components/WordList";
import { FetchWordData, fetchWords, Word } from "../data/fetchWords";
import { useAuth } from "../store/auth";

const Edit: React.FC = () => {
  const [createIsOpen, setCreateIsOpen] = useState(false);
  const [editIsOpen, setEditIsOpen] = useState(false);
  const [selectedWord, setSelectedWord] = useState<Word | undefined>(undefined);

  const idToken = useAuth((state) => state.idToken);

  const { isLoading, isError, data, error } = useQuery<FetchWordData, Error>(
    ["words", { isFibo: false, page: 1, limit: 10, idToken }],
    fetchWords,
    { staleTime: 3000 }
  );

  const wordList = data?.words ? (
    <WordList
      words={data.words}
      onWordSelected={(word) => {
        setSelectedWord(word);
        setEditIsOpen(true);
      }}
    />
  ) : undefined;

  return (
    <>
      <h1 className="title is-3">Your Word List</h1>
      <div className="buttons is-centered">
        <button
          onClick={() => setCreateIsOpen(true)}
          className="button is-info"
        >
          Create Word
        </button>
      </div>

      {isLoading && <Loader radius={200} />}
      {isError && <p>{error?.message}</p>}
      {wordList}

      {/* For creating a word */}
      <EditWordForm
        isOpen={createIsOpen}
        onClose={() => {
          setCreateIsOpen(false);
        }}
        onFormSubmitted={(values) => {
          console.log(values);
        }}
      />

      {/* For editing a word */}
      <EditWordForm
        isOpen={editIsOpen}
        initialWord={selectedWord}
        onClose={() => {
          setEditIsOpen(false);
          setSelectedWord(undefined);
        }}
        onFormSubmitted={(values) => {
          console.log(values);
        }}
      />
    </>
  );
};

export default Edit;
