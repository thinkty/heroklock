import React from 'react';

export const Footer = (): JSX.Element => {
  return (
    <div
      style={{
        position: 'absolute',
        bottom: 10,
        width: '100vw',
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'center',
        alignItems: 'center',
      }}
    >
      Based on &nbsp; <a href="https://kaffeine.herokuapp.com/" target="_blank">Kaffeine</a>
    </div>
  );
}
