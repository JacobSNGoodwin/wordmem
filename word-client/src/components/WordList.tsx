import React from "react";
import { Word } from "../data/fetchWords";
import WordListItem from "./WordListItem";

type WordListProps = {
  words: Word[];
  onWordSelected(word: Word): void;
};

const WordList: React.FC<WordListProps> = ({ words, onWordSelected }) => {
  const wordList = words.map((word) => (
    <div
      key={word.id}
      onClick={() => onWordSelected(word)}
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
