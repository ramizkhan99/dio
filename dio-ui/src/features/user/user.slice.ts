import { UserRegistrationRequest, UserRegistrationResponse } from '@components/registration/registration.types';
import { PayloadAction, createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { RegistrationAPI } from '../../components/registration/registration.api';
import { UserState, defaultState } from './user.state';

const userSlice = createSlice({
  name: 'user',
  initialState: defaultState,
  reducers: {
    setFirstName: (state, action: PayloadAction<string>) => {
      state.firstName = action.payload;
    },
    setLastName: (state, action: PayloadAction<string>) => {
      state.lastName = action.payload;
    },
    setEmailId: (state, action: PayloadAction<string>) => {
      state.emailId = action.payload;
    },
    setUsername: (state, action: PayloadAction<string>) => {
      state.username = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(registerUser.pending, (state: UserState, _) => {
      state.isPending = true;
    });
    builder.addCase(registerUser.fulfilled, (state: UserState, action) => {
      state.firstName = action.payload.firstName;
      state.lastName = action.payload.lastName;
      state.emailId = action.payload.emailId;
      state.username = action.payload.username;
      state.isPending = false;
    });
    builder.addCase(registerUser.rejected, (state, _) => {
      state.hasError = true;
      state.isPending = false;
    });
  },
});

export const registerUser = createAsyncThunk<UserRegistrationResponse, UserRegistrationRequest>(
  'user/register',
  async (request: UserRegistrationRequest, _) => {
    const registrationAPI: RegistrationAPI = new RegistrationAPI();
    const response = await registrationAPI.register(request);
    return response.data;
  }
);

export const { setFirstName } = userSlice.actions;

export const userReducer = userSlice.reducer;