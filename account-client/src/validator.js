import Vue from "vue";
import { ValidationProvider, ValidationObserver, extend } from "vee-validate";
import { required, email, min, max, confirmed } from "vee-validate/dist/rules";

extend("required", {
  ...required,
  message: "Field is required"
});
extend("email", {
  ...email,
  message: "Not a valid email address"
});
extend("min", {
  ...min,
  message: "Must be at least {length} character"
});
extend("max", {
  ...max,
  message: "Must be at maximum {length} character"
});
extend("confirmed", {
  ...confirmed,
  message: "Field must match {target}"
});

// custom rule to validate urls
extend("url", value => {
  const pattern = new RegExp(
    "^(https?:\\/\\/)" + // protocol
    "((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|" + // domain name
    "((\\d{1,3}\\.){3}\\d{1,3}))" + // OR ip (v4) address
    "(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*" + // port and path
    "(\\?[;&a-z\\d%_.~+=-]*)?" + // query string
      "(\\#[-a-z\\d_]*)?$",
    "i"
  ); // fragment locator
  return pattern.test(value) ? true : "{_field_} must be a valid url";
});

Vue.component("ValidationProvider", ValidationProvider);
Vue.component("ValidationObserver", ValidationObserver);
