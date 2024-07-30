import RegistrationFormComponent from '@components/registration/registration.component';
import React, { Component } from 'react';
import './registration.style.scss';

export class RegistrationPage extends Component {
  public render(): React.ReactNode {
    return (
      <div className='flex h-screen'>
        <RegistrationFormComponent />
      </div>
    );
  }
}