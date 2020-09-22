import React from "react";
import { Word } from "../data/fetchWords";

type WordCardProps = Word;

const WordCard: React.FC<WordCardProps> = (props) => {
  const startDate = new Date(props.startDate);
  const today = new Date();
  const msPerDay = 24 * 60 * 60 * 1000;

  const msDiff = today.getTime() - startDate.getTime();

  const days = Math.floor(msDiff / msPerDay);

  return (
    <div className="box">
      <div className="content">
        <div style={{ display: "flex" }}>
          <h4 className="mb-0">{props.word}</h4>
          <span className="tag is-link ml-4">
            {days === 1 ? "1 Day" : `${days} Days`}
          </span>
        </div>
        {props.refUrl ? (
          <a href={props.refUrl} target="_blank" rel="noopener noreferrer">
            <p className="help">Reference</p>
          </a>
        ) : undefined}
        <p className="mt-4">{props.definition}</p>
      </div>
    </div>
  );
};

export default WordCard;
