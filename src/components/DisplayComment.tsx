import { Box, Button, Typography } from "@mui/material";
import { Comment } from "../types/Comment";
import GetAvatar from "./GetAvatar";
import EditComment from "./EditComment";

interface Props {
  comment: Comment;
}

const DisplayComment: React.FC<Props> = ({ comment }) => {
  return (
    <Box
      sx={{
        border: 5,
        borderColor: "darkblue",
        padding: "15px",
        margin: "20px",
      }}
    >
      <GetAvatar user={comment.user} />
      <Typography variant="h5">{comment.user.name}</Typography>
      <Typography variant="body1">{comment.content}</Typography>
      <Button onClick={() => EditComment(comment)}>Edit</Button>
    </Box>
  );
};

export default DisplayComment;
