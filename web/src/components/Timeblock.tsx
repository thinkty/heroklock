import React from 'react';

export type TimeblockType = {
  URL: string;
  Duration: number;
  StartTime: Date;
  Valid: boolean;
};

type Props = TimeblockType & {
  delete: (url: string) => boolean;
} & typeof defaultProps;

const defaultProps = {};

export const Timeblock = (props: Props): JSX.Element => {
  return (
    <div
      style={{
        width: '100%',
        display: 'flex',
        flexDirection: 'row',
        
      }}
    >
      <h1>
        { props.URL }
      </h1>
    </div>
  );
}
Timeblock.defaultProps = defaultProps;
