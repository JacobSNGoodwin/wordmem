import { useFormik } from "formik";
import React, { useState } from "react";
import { useMutation, useQueryCache } from "react-query";
import updateWord from "../data/updateWord";
import { Word } from "../data/fetchWords";
import { useAuth } from "../store/auth";
import * as Yup from "yup";
import deleteWord from "../data/deleteWord";

type EditWordFormProps = {
  isOpen: boolean;
  initialWord?: Word;
  onClose(): void;
};

const EditWordForm: React.FC<EditWordFormProps> = ({
  isOpen,
  initialWord,
  onClose,
}) => {
  const { idToken } = useAuth();
  const queryCache = useQueryCache();
  const [errorMessage, setErrorMessage] = useState<string | undefined>(
    undefined
  );
  const [mutateUpdate, { isLoading: isUpdating }] = useMutation(updateWord, {
    onSuccess: async () => {
      setErrorMessage(undefined);
      queryCache.invalidateQueries("words");
      onClose();
    },
    onError: async (error: Error) => {
      setErrorMessage(error.message);
    },
  });
  const [mutateDelete, { isLoading: isDeleting }] = useMutation(deleteWord, {
    onSuccess: async () => {
      setErrorMessage(undefined);
      queryCache.invalidateQueries("words");
      onClose();
    },
    onError: async (error: Error) => {
      setErrorMessage(error.message);
    },
  });

  const formik = useFormik({
    initialValues: {
      word: initialWord?.word || "",
      definition: initialWord?.definition || "",
      refUrl: initialWord?.refUrl || "",
      startDate: initialWord?.startDate.substr(0, 10) || "", // substring gets YYYY-MM-DD out of date string
    },
    validationSchema: Yup.object({
      word: Yup.string().required("A non-empty word is required"),
      definition: Yup.string().required("A non-empty definition is required"),
      refUrl: Yup.string().url("Must be a valid URL or empty"),
      startDate: Yup.date(),
    }),
    onSubmit: (values) => {
      mutateUpdate({ ...values, id: initialWord?.id, idToken });
    },
    enableReinitialize: true, // doesn't populate fields on initial render
  });

  return (
    <div className={`modal${isOpen ? " is-active" : ""}`}>
      <div className="modal-background"></div>
      <div className="modal-card">
        <header className="modal-card-head">
          <p className="modal-card-title">
            {initialWord ? "Update Word" : "Create Word"}
          </p>
          <button
            onClick={() => {
              formik.resetForm();
              onClose();
            }}
            className="delete"
            aria-label="close"
          ></button>
        </header>
        <form onSubmit={formik.handleSubmit}>
          <section className="modal-card-body">
            <div className="field">
              <label htmlFor="word" className="label">
                Word
              </label>
              <div className="control">
                <input
                  id="word"
                  name="word"
                  className="input is-rounded"
                  type="text"
                  onChange={formik.handleChange}
                  onBlur={formik.handleBlur}
                  value={formik.values.word}
                />
              </div>
              {formik.touched.word && formik.errors.word && (
                <p className="has-text-centered has-text-danger">
                  {formik.errors.word}
                </p>
              )}
            </div>
            <div className="field">
              <label className="label">Definition</label>
              <div className="control">
                <textarea
                  id="definition"
                  name="definition"
                  className="textarea"
                  rows={4}
                  onChange={formik.handleChange}
                  onBlur={formik.handleBlur}
                  value={formik.values.definition}
                />
              </div>
              {formik.touched.definition && formik.errors.definition && (
                <p className="has-text-centered has-text-danger">
                  {formik.errors.definition}
                </p>
              )}
            </div>
            <div className="field">
              <label className="label">Reference URL</label>
              <div className="control">
                <input
                  id="refUrl"
                  name="refUrl"
                  className="input is-rounded"
                  type="text"
                  onChange={formik.handleChange}
                  onBlur={formik.handleBlur}
                  value={formik.values.refUrl}
                />
              </div>
              {formik.touched.refUrl && formik.errors.refUrl && (
                <p className="has-text-centered has-text-danger">
                  {formik.errors.refUrl}
                </p>
              )}
            </div>
            {initialWord && (
              <div className="field">
                <label className="label">Change Start Date</label>
                <div className="control">
                  <input
                    id="startDate"
                    name="startDate"
                    className="input is-rounded"
                    type="date"
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    value={formik.values.startDate}
                  />
                </div>
                {formik.touched.startDate && formik.errors.startDate && (
                  <p className="has-text-centered has-text-danger">
                    {formik.errors.startDate}
                  </p>
                )}
              </div>
            )}
            {errorMessage && (
              <p className="has-text-danger has-text-centered is-5">
                {errorMessage}
              </p>
            )}
          </section>
          <footer
            className="modal-card-foot"
            style={{ justifyContent: "space-between" }}
          >
            <button
              type="submit"
              className={`button is-info${isUpdating ? " is-loading" : ""}`}
              disabled={!formik.isValid || !formik.dirty}
            >
              Save changes
            </button>
            {initialWord && (
              <button
                onClick={() => mutateDelete({ id: initialWord.id, idToken })}
                type="button"
                className={`button is-danger${isDeleting ? " is-loading" : ""}`}
              >
                Delete?
              </button>
            )}
          </footer>
        </form>
      </div>
    </div>
  );
};

export default EditWordForm;
