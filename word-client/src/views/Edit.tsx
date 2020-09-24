import React, { useState } from "react";
import { useInfiniteQuery } from "react-query";
import EditWordForm from "../components/EditWordForm";
import Loader from "../components/ui/Loader";
import WordList from "../components/WordList";
import InfiniteScroll from "react-infinite-scroller";
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
    data.map((group, i) => (
      <React.Fragment key={i}>
        {group.words && (
          <WordList
            key={i}
            words={group.words}
            onWordSelected={(word) => {
              setSelectedWord(word);
              setEditIsOpen(true);
            }}
          />
        )}
      </React.Fragment>
    ));

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
      {wordList && (
        <InfiniteScroll
          className="columns is-multiline mt-6"
          pageStart={0}
          loadMore={() => fetchMore()}
          hasMore={canFetchMore}
          loader={<Loader key={0} color="red" />}
          children={wordList}
        />
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
