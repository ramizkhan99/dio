import { SyntheticEvent } from 'react';

export interface TextInputProps {
  /**
   * Type of the input field (text/password/etc.)
   */
  type: string;

  /**
   * Placeholder text of the input field
   */
  placeholder?: string;

  /**
   * Text of the label field (Not including would remove label field)
   */
  labelText?: string;

  /**
   * Name of the input field
   */
  name: string;

  /**
   * Additional classes
   */
  class?: string;

  /**
   * Register function on change in input field data
   */
  onChange?(event: SyntheticEvent): void;
}