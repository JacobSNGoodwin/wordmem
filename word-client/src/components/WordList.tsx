import React from "react";
import { Word } from "../data/fetchWords";
import WordListItem from "./WordListItem";

type WordListProps = {
  words: Word[];
};

const WordList: React.FC<WordListProps> = ({ words }) => {
  const wordList = words.map((word) => (
    <div
      key={word.id}
      className="column is-half-tablet is-one-third-desktop is-one-quarter-widescreen"
      style={{ cursor: "pointer" }}
    >
      <WordListItem {...word} />
    </div>
  ));
  return (
    <div className="mt-6">
      <div className="columns is-multiline">{wordList}</div>
    </div>
  );
};

export default WordList;
