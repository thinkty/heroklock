import React from 'react';

type Props = {
  onSubmit: (url: string, dur: number) => void;
} & typeof defaultProps;

const defaultProps = {};

export const Form = (props: Props): JSX.Element => {
  const [url, setUrl] = React.useState<string>("");
  const [dur, setDur] = React.useState<number>(0);

  return (
    <form
      onSubmit={(e) => { e.preventDefault(); props.onSubmit(url, dur); setUrl(""); setDur(0); }}
      style={{
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        gap: 20,
      }}
    >
      {/* URL input field */}
      <div
        style={{
          display: 'flex',
          flexDirection: 'row',
          gap: 0,
        }}
      >
        <label htmlFor="url">https://</label>
        <input
          required
          name="url"
          type="text"
          value={url}
          minLength={1}
          maxLength={30}
          onChange={(e) => { setUrl(e.target.value) }}
          style={{
            color: 'white',
            backgroundColor: 'transparent',
            border: 'none',
            borderBottom: 'thin solid white',
            outline: 'none',
            textAlign: 'center',
          }}
        />
        <label htmlFor="url">.herokuapp.com</label>
      </div>

      {/* Duration input field */}
      <div
        style={{
          display: 'flex',
          flexDirection: 'row',
          gap: 20,
        }}
      >
        <input
          required
          name="duration"
          type="number"
          value={dur}
          min={1}
          max={1460}
          onChange={(e) => { setDur(parseInt(e.target.value)) }}
          style={{
            width: '50%',
            color: 'white',
            backgroundColor: 'transparent',
            border: 'none',
            borderBottom: 'thin solid white',
            outline: 'none',
            textAlign: 'center',
          }}
        />
        <label htmlFor="duration" style={{ width: '50%' }}>
          { dur / 2 + " Hours" }
        </label>
      </div>

      <input
        type="submit"
        value="Submit"
        style={{
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
          border: 'thin solid white',
          borderRadius: 10,
          padding: 10,
        }}
      />
    </form>
  );
}
Form.defaultProps = defaultProps;
