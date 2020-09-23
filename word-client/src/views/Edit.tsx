import React, { useState } from "react";
import { useInfiniteQuery } from "react-query";
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

  const {
    data,
    isLoading,
    error,
    isError,
    canFetchMore,
    fetchMore,
    isFetchingMore,
  } = useInfiniteQuery<FetchWordData, Error>(
    ["words", { isFibo: false, limit: 4, idToken }],
    fetchWords,
    {
      getFetchMore: (lastGroup, allGroups) => {
        // this function returns query values for next query
        const { page, pages } = lastGroup;

        // if there are no more queries, returns undefined
        if (page >= pages) {
          return undefined;
        }

        return page + 1;
      },
    }
  );

  const wordList =
    data &&
    data.map(
      (group, i) =>
        group.words && (
          <WordList
            key={i}
            words={group.words}
            onWordSelected={(word) => {
              setSelectedWord(word);
              setEditIsOpen(true);
            }}
          />
        )
    );

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
      {isFetchingMore && <Loader color="red" />}
      {canFetchMore && (
        <button
          onClick={() => {
            fetchMore();
          }} // do not pass reference, or else event gets sent as argument to fetchMore!
          className="button is-primary mt-6"
        >
          Fetch more!
        </button>
      )}

      {/* For creating a word */}
      <EditWordForm
        isOpen={createIsOpen}
        onClose={() => {
          setCreateIsOpen(false);
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
      />
    </>
  );
};

export default Edit;
