import { TextInput } from '@components/text-input';
import { debounce } from 'lodash';
import { Component, ReactNode, SyntheticEvent } from 'react';
import { connect } from 'react-redux';
import { registerUser } from '../../features/user/user.slice';
import { UserState } from '../../features/user/user.state';
import './registration.style.scss';
import { RegistrationProps, UserRegistrationRequest } from './registration.types';


export class RegistrationFormComponent extends Component<RegistrationProps> {

  private userRequest: UserRegistrationRequest = {
    firstName: '',
    lastName: '',
    username: '',
    emailId: '',
    password: ''
  };

  private async handleSubmit(event: SyntheticEvent): Promise<void> {
    event.preventDefault();

    this.props.registerUser({
      ...this.userRequest,
      password: 'random'
    });
  }

  private handleFirstNameChange(event: SyntheticEvent): void {
    this.userRequest.firstName = (event.target as HTMLInputElement).value;
  }

  private handleLastNameChange(event: SyntheticEvent): void {
    this.userRequest.lastName = (event.target as HTMLInputElement).value;
  }

  private handleUsernameChange(event: SyntheticEvent): void {
    this.userRequest.username = (event.target as HTMLInputElement).value;
  }

  private handleEmailidChange(event: SyntheticEvent): void {
    this.userRequest.emailId = (event.target as HTMLInputElement).value;
  }

  private handlePasswordChange(event: SyntheticEvent): void {
    this.userRequest.password = (event.target as HTMLInputElement).value;
  }

  public render(): ReactNode {
    return (
      <div className='container w-4/12 m-auto card px-4 registration-form'>
        <div className="card-title prose mt-2 pt-2 pb-2">
          <h2 className='w-full text-center'>Registration Form</h2>
        </div>
        <div className="card-body">
          <form className='form-control' onSubmit={this.handleSubmit.bind(this)}>
            <div className="flex flex-row">
              <TextInput
                type='input'
                placeholder='Firstname'
                labelText='Firstname'
                name='firstname'
                onChange={debounce(this.handleFirstNameChange.bind(this), 300)}
                class='w-full'
              />
              <TextInput
                type='input'
                placeholder='Lastname'
                labelText='Lastname'
                name='lastname'
                onChange={debounce(this.handleLastNameChange.bind(this))}
                class='w-full'
              />
            </div>
            <TextInput
              type='input'
              placeholder='Username'
              labelText='Username'
              name='username'
              onChange={debounce(this.handleUsernameChange.bind(this))}
            />
            <TextInput
              type='email'
              placeholder='Email'
              labelText='Email'
              name='email'
              onChange={debounce(this.handleEmailidChange.bind(this))}
            />
            <TextInput
              type='password'
              placeholder='Password'
              labelText='Password'
              name='password'
              onChange={debounce(this.handlePasswordChange.bind(this))}
            />
            <button
              className='btn btn-block btn-success mt-2 text-white' type='submit' onSubmit={this.handleSubmit}>
              Submit<i className='ri-arrow-right-s-line text-lg'></i>
            </button>
          </form>
        </div>
      </div>
    );
  }
}

const mapStateToProps = (state: UserState) => ({
  username: state.username,
  emailId: state.emailId,
  firstName: state.firstName,
  lastName: state.lastName,
  isPending: state.isPending,
});

const mapDispatchToProps = {
  registerUser,
};

export default connect(mapStateToProps, mapDispatchToProps)(RegistrationFormComponent);