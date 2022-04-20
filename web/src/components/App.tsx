import React from 'react';
import { Footer } from './Footer';
import { Form } from './Form';

type Props = {
  message: string;
} & typeof defaultProps;

const defaultProps = {
  message: 'Heroklock',
};

export const App = (props: Props): JSX.Element => {
  return (
    <div
      style={{
        width: '100vw',
        height: '100vh',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        fontFamily: 'Verdana, sans-serif',
        backgroundColor: 'black',
        color: 'white',
      }}
    >
      <h1>
        { props.message }
      </h1>
      <Form
        onSubmit={(url, dur) => { console.log(url, dur) }}
      />
      <Footer />
    </div>
  );
}
App.defaultProps = defaultProps;
