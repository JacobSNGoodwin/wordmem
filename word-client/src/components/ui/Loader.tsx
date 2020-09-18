import React from "react";
import styles from "./Loader.module.css";

type LoaderProps = {
  color?: string;
  radius?: number;
};

const Loader: React.FC<LoaderProps> = ({
  color = "#3273dc",
  radius = 80,
}: LoaderProps) => {
  return (
    <div
      className={styles.spinner}
      style={{ width: `${radius}px`, height: `${radius}px` }}
    >
      <div
        className={styles["double-bounce1"]}
        style={{ backgroundColor: color }}
      ></div>
      <div
        className={styles["double-bounce2"]}
        style={{ backgroundColor: color }}
      ></div>
    </div>
  );
};

export default Loader;
