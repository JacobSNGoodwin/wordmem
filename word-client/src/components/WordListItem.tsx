import React from "react";
import { Word } from "../data/fetchWords";
import { daysSinceCreation } from "../util";

const WordListItem: React.FC<Word> = (props) => {
  const creationDate = new Date(props.startDate);
  const daysSinceCreated = daysSinceCreation(props.startDate);
  const year = creationDate.getFullYear();
  const month = creationDate.getMonth();
  const day = creationDate.getDate(); // day returns day of week, 0-6

  // using click to be able to have inner link inside of card
  const refUrl = props.refUrl ? (
    <>
      <div
        onClick={(event) => openRefUrl(event)}
        style={{ textDecoration: "underline" }}
      >
        Reference
      </div>
    </>
  ) : undefined;

  const openRefUrl = (event: React.MouseEvent<HTMLElement, MouseEvent>) => {
    event.stopPropagation();
    window.open(props.refUrl);
  };

  return (
    <div
      className="notification is-link"
      style={{
        height: "100%",
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
      }}
    >
      <div>
        <div className="is-capitalized has-text-weight-bold">{props.word}</div>
        <div className="is-italic">{props.definition}</div>
      </div>
      <div style={{ marginTop: 24 }}>
        <div>{refUrl}</div>
        <div className="has-text-weight-semibold">
          Creation Date: {`${year}-${month}-${day}`}
        </div>
        <div className="has-text-weight-semibold">
          Days since creation: {daysSinceCreated}
        </div>
      </div>
    </div>
  );
};

export default WordListItem;
