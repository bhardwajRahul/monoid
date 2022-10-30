import React from 'react';
import { Link } from 'react-router-dom';
import { classNames } from '../utils/utils';

interface ButtonProps extends Omit<React.HTMLProps<HTMLButtonElement>, 'size'> {
  size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl',
  variant?: 'primary' | 'white' | 'danger',
  type?: 'button' | 'link',
  to?: string
}

export default function Button(props: ButtonProps) {
  let classes = '';
  const {
    size, variant, className, children, type, to,
  } = props;

  switch (size) {
    case 'xs':
      classes = 'inline-flex items-center px-2.5 py-1.5 border text-xs font-medium rounded shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-2';
      break;
    case 'sm':
      classes = 'inline-flex items-center px-3 py-2 border text-sm leading-4 font-medium rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-2';
      break;
    case 'lg':
      classes = 'inline-flex items-center px-4 py-2 border text-base font-medium rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-2';
      break;
    case 'xl':
      classes = 'inline-flex items-center px-6 py-3 border text-base font-medium rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-2';
      break;
    default:
      classes = 'inline-flex items-center px-4 py-2 border text-sm font-medium rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-offset-2';
      break;
  }

  switch (variant) {
    case 'white':
      classes = classNames(classes, 'border-px border-gray-300 bg-white text-gray-500 hover:bg-gray-50 shadow-sm');
      break;
    case 'danger':
      classes = classNames(classes, 'text-white bg-red-600 hover:bg-red-700 focus:ring-red-500 border-transparent');
      break;
    default:
      classes = classNames(classes, 'text-white bg-indigo-600 hover:bg-indigo-700 focus:ring-indigo-500 border-transparent');
  }

  classes = classNames(classes, className);

  if (type === 'link') {
    let toStr = to;

    if (toStr === undefined) {
      toStr = '#';
    }

    return (
      <Link
        className={classes}
        to={toStr}
      >
        {children}
      </Link>
    );
  }

  return (
    <button
      {...props}
      type="button"
      className={classes}
    >
      {children}
    </button>
  );
}

Button.defaultProps = {
  size: 'md',
  variant: 'primary',
  type: 'button',
  to: undefined,
};