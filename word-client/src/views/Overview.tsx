import React from "react";
import InfiniteScroll from "react-infinite-scroller";
import { useInfiniteQuery } from "react-query";
import Loader from "../components/ui/Loader";
import WordCard from "../components/WordCard";
import { FetchWordData, fetchWords } from "../data/fetchWords";
import { useAuth } from "../store/auth";

const Overview: React.FC = () => {
  const idToken = useAuth((state) => state.idToken);

  const { data, isLoading, error, canFetchMore, fetchMore } = useInfiniteQuery<
    FetchWordData,
    Error
  >(["words", { isFibo: true, limit: 4, idToken }], fetchWords, {
    getFetchMore: (lastGroup, allGroups) => {
      // this function returns query values for next query
      const { page, pages } = lastGroup;

      // if there are no more queries, returns undefined
      if (page >= pages) {
        return undefined;
      }

      return page + 1;
    },
  });

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
      {wordList && (
        <InfiniteScroll
          pageStart={0}
          loadMore={() => fetchMore()}
          hasMore={canFetchMore}
          loader={<Loader key={0} color="red" />}
          children={wordList}
        />
      )}
    </>
  );
};

export default Overview;
