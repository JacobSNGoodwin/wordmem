import React from "react";
import { useInfiniteQuery } from "react-query";
import Loader from "../components/ui/Loader";
import WordCard from "../components/WordCard";
import { FetchWordData, fetchWords } from "../data/fetchWords";
import { useAuth } from "../store/auth";

const Overview: React.FC = () => {
  const idToken = useAuth((state) => state.idToken);

  const {
    data,
    isLoading,
    error,
    canFetchMore,
    fetchMore,
    isFetchingMore,
  } = useInfiniteQuery<FetchWordData, Error>(
    ["words", { isFibo: true, limit: 2, idToken }],
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
        {group.words.map((word) => (
          <WordCard key={word.id} {...word} />
        ))}
      </React.Fragment>
    ));
  return (
    <>
      <h1 className="title is-3">Today's Words</h1>

      {isLoading && <Loader radius={200} />}
      {error && <p>{error.message}</p>}
      {wordList}
      {isFetchingMore && <Loader color="red" />}
      {canFetchMore && (
        <button
          onClick={() => {
            fetchMore();
          }} // do not pass reference, or else event gets sent as argument to fetchMore!
          className="button is-primary"
        >
          Fetch more!
        </button>
      )}
    </>
  );
};

export default Overview;
