<template>
  <div class="my-6">
    <!-- Image stuffs -->
    <figure class="image is-128x128">
      <div class="placeholder">
        <img v-if="imageUrl" :src="imageUrl" />
      </div>
    </figure>
    <div class="buttons is-centered my-6">
      <button class="button is-link is-rounded" @click="openImageSelector">
        Update Image
      </button>
      <button
        @click="deleteUserImage"
        class="button is-danger is-rounded"
        :disabled="!imageUrl"
        :class="{ 'is-loading': isDeleteing }"
      >
        Delete Image
      </button>
    </div>
    <div v-if="deleteError">{{ deleteError }}</div>
    <ImageSelector
      :isActive="imageSelectorActive"
      @imageUrlUpdated="updateImageUrl"
      @close="closeImageSelector"
    />

    <!-- Details Form -->
    <div class="columns is-centered">
      <div class="column is-half-desktop">
        <ValidationObserver v-slot="{ handleSubmit, invalid, changed }">
          <form
            novalidate="true"
            @submit.prevent="handleSubmit(submitUserDetails)"
          >
            <div class="field my-5">
              <div class="control">
                <ValidationProvider name="email" rules="required" v-slot="v">
                  <input
                    class="input is-rounded has-text-weight-bold is-centered"
                    type="email"
                    v-model="email"
                    placeholder="Email Address"
                  />
                  <div
                    v-if="v.touched && v.invalid"
                    class="help is-danger has-text-centered"
                  >
                    <p v-for="error in v.errors" :key="error">
                      {{ error }}
                    </p>
                  </div>
                </ValidationProvider>
              </div>
            </div>
            <!-- In reality I wouldn't validate a person's name so stringently,
                 but for the sake of showing another validator, why not, eh! 
             -->
            <div class="field mb-5">
              <div class="control">
                <ValidationProvider name="name" rules="max:40" v-slot="v">
                  <input
                    class="input is-rounded has-text-weight-bold is-centered"
                    type="text"
                    v-model="name"
                    placeholder="Name"
                  />
                  <div
                    v-if="v.touched && v.invalid"
                    class="help is-danger has-text-centered"
                  >
                    <p v-for="error in v.errors" :key="error">
                      {{ error }}
                    </p>
                  </div>
                </ValidationProvider>
              </div>
            </div>
            <div class="field mb-5">
              <div class="control">
                <ValidationProvider name="website" rules="url" v-slot="v">
                  <input
                    class="input is-rounded has-text-weight-bold is-centered"
                    type="text"
                    v-model="website"
                    placeholder="Wesbite URL"
                  />
                  <div
                    v-if="v.touched && v.invalid"
                    class="help is-danger has-text-centered"
                  >
                    <p v-for="error in v.errors" :key="error">
                      {{ error }}
                    </p>
                  </div>
                </ValidationProvider>
              </div>
            </div>

            <div class="buttons is-centered mt-6">
              <button
                type="submit"
                :disabled="invalid || !changed"
                class="button is-info is-rounded"
                :class="{ 'is-loading': isUpdating }"
              >
                Update Details
              </button>
            </div>
          </form>
        </ValidationObserver>
      </div>
    </div>
  </div>
</template>

<script>
import ImageSelector from "./ImageSelector";
import useRequest from "../composables/useRequest";
import { useAuth } from "../store/auth";

export default {
  name: "UpdateForm",
  components: {
    ImageSelector
  },
  props: {
    user: {
      type: Object,
      default: null
    }
  },
  setup() {
    const { idToken } = useAuth();

    const {
      deleteImage,
      error: deleteError,
      loading: isDeleteing
    } = useRequest({
      url: "/api/image",
      method: "delete",
      headers: {
        Authorization: `Bearer ${idToken.value}`
      }
    });

    const {
      exec: updateUser,
      data: updateUserData,
      error: updateError,
      loading: isUpdating
    } = useRequest({
      url: "/api/details",
      method: "put",
      headers: {
        Authorization: `Bearer ${idToken.value}`
      }
    });

    return {
      deleteImage,
      deleteError,
      isDeleteing,
      updateUser,
      updateUserData,
      updateError,
      isUpdating
    };
  },
  data() {
    return {
      name: this.user.name,
      email: this.user.email,
      website: this.user.website,
      imageUrl: this.user.imageUrl,
      imageSelectorActive: false
    };
  },
  methods: {
    openImageSelector() {
      this.imageSelectorActive = true;
    },
    closeImageSelector() {
      this.imageSelectorActive = false;
    },
    updateImageUrl(newImageUrl) {
      this.imageUrl = newImageUrl + `?_${Date.now()}`; // trick to refresh cached value
      this.imageSelectorActive = false;
    },
    async deleteUserImage() {
      // probably this should be built into the composable
      // as we might have to repeat this callback in other components
      await this.deleteImage();

      if (!this.deleteError) {
        this.imageUrl = "";
      }
    },
    submitUserDetails() {
      const data = {
        name: this.name,
        email: this.email,
        website: this.website
      };
      this.updateUser(data);
    }
  },
  watch: {
    updateUserData: function({ user }) {
      this.name = user.name;
      this.email = user.email;
      this.website = user.website;
    }
  }
};
</script>

<style scoped lang="scss">
.image {
  margin: auto;
}

.placeholder {
  width: 128px;
  height: 128px;
  background-color: #f5f5f5;
  border-radius: 64px;

  img {
    height: 128px;
    border-radius: 64px;
  }
}

button {
  width: 120px;
}
</style>
