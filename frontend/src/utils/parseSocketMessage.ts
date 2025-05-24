// utils/parseSocketMessage.ts
export function parseSocketMessage(data: string) {
  const parsed = JSON.parse(data)
  return {
    message: parsed.message,
    inputConfig: {
      minLength: parsed.minLength,
      maxLength: parsed.maxLength,
      reference: parsed.reference,
      required: parsed.required,
      type: parsed.type,
    }
  }
}