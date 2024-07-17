type Type =
  | "text"
  | "contacts"
  | "audio"
  | "document"
  | "image"
  | "interactive"
  | "location"
  | "reaction"
  | "sticker"
  | "video";

export class ResponsePayload {
  messaging_product: string;
  contacts: {
    input: string;
    wa_id: string;
  }[];
  messages: { id: string }[];
}

export class RequestPayload {
  appToken: string;
  phoneNumberId: string;
  data:
    | TextModel
    | ContactsModel
    | AudioModel
    | DocumentModel
    | ImageModel
    | InteractiveCtaUrlModel
    | InteractiveFlowModel
    | InteractiveListModel
    | InteractiveReplyButtonModel
    | LocationModel
    | LocationRequestModel
    | ReactionModel
    | StickerModel
    | VideoModel;
}

class HeaderImageModel {
  type: "image";
  image: {
    id?: string;
    link?: string;
  };
}

class HeaderDocumentModel {
  type: "document";
  document: {
    id?: string;
    link?: string;
  };
}

class HeaderVideoModel {
  type: "video";
  video: {
    id?: string;
    link?: string;
  };
}

class HeaderTextModel {
  type: "text";
  text: string;
}

class MessageData {
  messaging_product: "whatsapp";
  recipient_type?: "individual";
  to: string; // O ID do WhatsApp ou o número de telefone do cliente que deve receber a mensagem.
  type?: Type; // O tipo de mensagem.
}

export class TextModel extends MessageData {
  text: {
    preview_url?: boolean; // Visualização do link de qualquer URL na sequência de texto do corpo.
    body: string; // Obrigatório para mensagens de texto. Comprimento máximo: 4096 caracteres.
  };
}

export class ContactsModel extends MessageData {
  contacts: {
    addresses: {
      street?: string;
      city?: string;
      state?: string;
      zip?: string;
      country?: string;
      country_code?: string;
      type?: string;
    }[];
    birthday?: string;
    emails: {
      email?: string;
      type?: string;
    }[];
    name: {
      formatted_name: string;
      first_name?: string;
      last_name?: string;
      middle_name?: string;
      suffix?: string;
      prefix?: string;
    };
    org: {
      company?: string;
      department?: string;
      title?: string;
    };
    phones: {
      phone?: string;
      type?: string;
      wa_id?: string;
    }[];
    urls: {
      url?: string;
      type?: string;
    }[];
  }[];
}

export class AudioModel extends MessageData {
  // Tipos (.acc, .amr, .mp3, .m4a, .ogg) max 16mb
  audio: {
    id?: string;
    link?: string;
  };
}

export class DocumentModel extends MessageData {
  // tipos (.txt, .xls, .xlsx, .doc, .docx, .ppt, .pptx, .pdf) max size 100mb
  document: {
    id?: string /* Only if using uploaded media */;
    link?: string /* Only if linking to your media */;
    caption?: string;
    filename?: string;
  };
}

export class ImageModel extends MessageData {
  // tipos (.jpeg, .png) max size 5mb
  image: {
    id?: string /* Only if using uploaded media */;
    link?: string /* Only if linking to your media */;
    caption?: string;
  };
}

export class InteractiveCtaUrlModel extends MessageData {
  interactive: {
    type: "cta_url";
    header?: {
      type: "text";
      text: string;
    };
    body?: {
      text: string;
    };
    footer?: {
      text: string;
    };
    action: {
      name: "cta_url";
      parameters: {
        display_text: string;
        url: string;
      };
    };
  };
}

export class InteractiveFlowModel extends MessageData {
  // doc: https://developers.facebook.com/docs/whatsapp/cloud-api/messages/interactive-flow-messages
  interactive: {
    type: "flow";
    header?: {
      type: "text";
      text: string;
    };
    body?: {
      text: string;
    };
    footer?: {
      text: string;
    };
    action: {
      name: "flow";
      parameters: {
        flow_message_version: string;
        flow_token: string;
        flow_id: string;
        flow_cta: string;
        flow_action: string;
        flow_action_payload: {
          screen: string;
          data: {
            product_name: string;
            product_description: string;
            product_price: number;
          };
        };
      };
    };
  };
}

export class InteractiveListModel extends MessageData {
  interactive: {
    type: "list";
    header?: {
      type: "text";
      text: string; // Supports text header type only. Maximum 60 characters.
    };
    body?: {
      text: string; // Supports URLs. Maximum 4096 characters.
    };
    footer?: {
      text: string; // Maximum 60 characters.
    };
    action: {
      sections: {
        title: string; // É necessária pelo menos 1 seção. Suporta até 10 seções. Máximo de 24 caracteres.
        rows: {
          id: string; // String arbitrária que identifica a linha. Esse ID será incluído na carga útil do webhook se o usuário enviar a seleção. É necessária pelo menos uma linha. Suporta até 10 linhas. Máximo de 200 caracteres.
          title: string; // Maximum 24 characters.
          description?: string; // Maximum 72 characters.
        }[];
      }[];
      button: string; // Button label text. When tapped, reveals rows (options the WhatsApp user can tap). Supports a single button. Maximum 20 characters.
    };
  };
}

export class InteractiveReplyButtonModel extends MessageData {
  interactive: {
    type: "button";
    header?:
      | HeaderImageModel
      | HeaderDocumentModel
      | HeaderVideoModel
      | HeaderTextModel;
    body: {
      text: string; // Emojis, descontos e links são suportados. Máximo de 1024 caracteres.
    };
    footer?: {
      text: string; // Emojis, markdown, and links are supported. Maximum 60 characters.
    };
    action: {
      buttons: {
        type: "reply";
        reply: {
          id: string; // A unique identifier for each button. Supports up to 3 buttons. Maximum 256 characters.
          title: string; // Maximum 20 characters.
        };
      }[];
    };
  };
}

export class LocationModel extends MessageData {
  location: {
    latitude: string;
    longitude: string;
    name?: string;
    address?: string;
  };
}

export class LocationRequestModel extends MessageData {
  interactive: {
    type: "location_request_message";
    body: {
      text: string;
    };
    action: {
      name: string;
    };
  };
}

export class ReactionModel extends MessageData {
  reaction: {
    message_id: string;
    emoji: string; // Sequência de escape Unicode do emoji, ou do próprio emoji, para aplicar à mensagem do usuário.
  };
}

export class StickerModel extends MessageData {
  // tipos (.webp, .webp) max size 100kb
  sticker: {
    id?: string;
    link?: string;
  };
}

export class VideoModel extends MessageData {
  // tipos (.3gp, .mp4) max size 16mb
  video: {
    id?: string;
    link?: string;
    caption?: string; // Maximum 1024 characters.
  };
}

export class ReadMessageModel {
  messaging_product: "whatsapp";
  status: "read";
  message_id: string;
}

// como fazer upload de media para enviar por ID https://developers.facebook.com/docs/whatsapp/cloud-api/reference/media#upload-media
