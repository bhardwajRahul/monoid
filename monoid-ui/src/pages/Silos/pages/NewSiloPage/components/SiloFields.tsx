import { gql, useQuery } from '@apollo/client';
import React, { useMemo } from 'react';
import Spinner from '../../../../../components/Spinner';
import { SiloSpec } from '../../../../../lib/models';
import Input, { InputLabel } from '../../../../../components/Input';
import { MonoidJSONSchema } from '../../../../../lib/types';
import Toggle from '../../../../../components/Toggle';
import TextMultiInput from '../../../../../components/TextMultiInput';
import BorderedRegion from '../../../../../components/BorderedRegion';

const GET_SILO_DATA = gql`
  query GetSiloSpecs($id: ID!) {
    siloSpecification(id: $id) {
      id
      schema
    }
  }
`;

const cmp = (v1?: number, v2?: number) => {
  const v1c = v1 === undefined ? 1000 : v1;
  const v2c = v2 === undefined ? 1000 : v2;

  return v1c - v2c;
};

function JSONSchemaControl(
  props: {
    def: MonoidJSONSchema,
    value: any,
    onChange: (v: any) => void,
    root?: boolean
  },
) {
  const {
    def, value, onChange, root,
  } = props;

  switch (def.type) {
    case 'string':
      return (
        <Input
          value={value}
          onChange={(e) => {
            onChange(e.target.value);
          }}
          type={def.secret ? 'password' : undefined}
        />
      );
    case 'integer':
      return (
        <Input
          value={value}
          onChange={(e) => {
            onChange(parseInt(e.target.value, 10));
          }}
          type="number"
        />
      );
    case 'number':
      return (
        <Input
          value={value}
          onChange={(e) => {
            onChange(parseFloat(e.target.value));
          }}
          type="number"
        />
      );
    case 'boolean':
      return (
        <Toggle
          checked={value}
          onChange={(v) => {
            onChange(v);
          }}
          size="lg"
        />
      );
    case 'array':
      switch ((def.items as MonoidJSONSchema).type) {
        case 'string':
          return (
            <TextMultiInput
              value={value ? value as string[] : []}
              onChange={(n) => onChange(n)}
              placeholder=""
            />
          );
        default:
          return <div />;
          break;
      }
      break;
    case 'object': {
      const fields = (
        Object.keys(def.properties!).map((k) => (
          { k, v: def.properties![k] as MonoidJSONSchema }
        )).sort((o1, o2) => cmp(o1.v.order, o2.v.order)).map(({ k, v }) => (
          // eslint-disable-next-line @typescript-eslint/no-use-before-define
          <JSONSchemaField
            def={v as MonoidJSONSchema}
            value={value ? value[k] : undefined}
            onChange={(val) => {
              onChange({
                ...value,
                [k]: val,
              });
            }}
            key={k}
          />
        ))
      );

      if (root) {
        // eslint-disable-next-line react/jsx-no-useless-fragment
        return <>{fields}</>;
      }
      return (
        <BorderedRegion label={def.title || 'No Name Provided'}>
          <div className="space-y-6">
            {fields}
          </div>
        </BorderedRegion>
      );
    }
    default:
      return (
        <div>
          Unsupported type
        </div>
      );
  }
}

JSONSchemaControl.defaultProps = {
  root: false,
};

const buildDefaultObject: (def: MonoidJSONSchema) => any = (def: MonoidJSONSchema) => {
  if (def.default) {
    return def.default;
  }

  switch (def.type) {
    case 'array':
      return [];
    case 'boolean':
      return false;
    case 'integer':
      return 0;
    case 'number':
      return 0;
    case 'object':
      return Object.fromEntries(
        Object.keys(def.properties!).map((k) => (
          { k, v: def.properties![k] as MonoidJSONSchema }
        )).map(({ k, v }) => (
          [k, buildDefaultObject(v)]
        )),
      );
    case 'string':
      return '';
    default:
      return null;
  }

  return null;
};

function JSONSchemaField(
  props: {
    def: MonoidJSONSchema,
    value: any,
    onChange: (v: any) => void
  },
) {
  const { def, value, onChange } = props;

  return (
    <div>
      <InputLabel htmlFor="siloName">
        {def.title}
      </InputLabel>
      <div className="mt-2">
        <JSONSchemaControl
          def={def}
          value={value}
          onChange={onChange}
        />
      </div>
    </div>
  );
}

export default function SiloFields(props: {
  siloID: string,
  siloData: any,
  setSiloData: (v: any) => void
}) {
  const { siloID, siloData, setSiloData } = props;
  const { data, loading, error } = useQuery<{ siloSpecification: SiloSpec }>(GET_SILO_DATA, {
    variables: {
      id: siloID,
    },
  });

  const jsonSchema = useMemo(() => {
    if (!data?.siloSpecification) {
      return undefined;
    }

    const schema = JSON.parse(data.siloSpecification.schema) as MonoidJSONSchema;
    setSiloData(buildDefaultObject(schema));
    return schema;
  }, [data?.siloSpecification]);

  if (error) {
    return (
      <div>{error.message}</div>
    );
  }

  if (loading || !jsonSchema) {
    return <Spinner />;
  }

  return (
    <JSONSchemaControl def={jsonSchema} value={siloData} onChange={(v) => setSiloData(v)} root />
  );
}