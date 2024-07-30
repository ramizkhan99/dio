import { Component, ReactNode } from 'react';
import { TextInputProps } from './text-input.props';

export class TextInput extends Component<TextInputProps> {
  public render(): ReactNode {
    return (
      <div className={`form-item m-2 ${this.props.class}`}>
        {this.props.labelText && (
          <label className='label'>
            <span className='label-text'>{this.props.labelText}</span>
          </label>
        )}
        <input
          type={this.props.type}
          placeholder={this.props.placeholder}
          className='input input-bordered w-full'
          onChange={this.props.onChange}
          name={this.props.name} />
      </div>
    );
  }
}